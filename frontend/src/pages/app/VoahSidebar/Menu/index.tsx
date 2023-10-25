import { atom, useAtom } from 'jotai';
import { KeyRoundIcon } from 'lucide-react';
import { NavigateFunction, useNavigate } from 'react-router-dom';

import { LucideCustom } from '@/lib/LucideCustom';

export interface Menu {
  icon: JSX.Element | string;
  name: string;
  onClick: ((e: React.MouseEvent<HTMLButtonElement>) => void) | string;
  mentioned?: number;
  isFocused?: boolean;
  categoryId?: string;
  subButton?: {
    icon: JSX.Element | string;
    onClick: ((e: React.MouseEvent<HTMLButtonElement>) => void) | string;
  };
  subMenu?: Array<Menu>;
}

export interface ParsedMenu extends Menu {
  icon: JSX.Element;
  onClick: (e: React.MouseEvent<HTMLButtonElement>) => void;
  subButton?: {
    icon: JSX.Element;
    onClick: (e: React.MouseEvent<HTMLButtonElement>) => void;
  };
  subMenu?: Array<ParsedMenu>;
}

export const menuAtom = atom<{
  info: {
    title: string;
    desc: string;
    hideDesc: boolean;
  };
  categories: Array<{
    id: string;
    name: string;
  }>;
  menus: Array<ParsedMenu>;
}>({
  info: {
    title: 'VOAH TITLE',
    desc: 'VOAH DESC',
    hideDesc: false,
  },
  categories: [],
  menus: [
    {
      icon: <KeyRoundIcon size={20} />,
      name: '개인정보 및 보안',
      onClick: () => {
        alert('준비중입니다.');
      },
      subMenu: [],
    },
  ],
});

function parseIcon(icon: JSX.Element | string): JSX.Element {
  if (typeof icon !== 'string') return icon;
  const temp = icon.split('::');
  return (
    <LucideCustom
      icon={temp[0]}
      size={Number(temp[1]) || 24}
      color={temp[2] || 'currentColor'}
    />
  );
}

function parseOnClick(
  onClick: ((e: React.MouseEvent<HTMLButtonElement>) => void) | string,
  navigate: NavigateFunction,
) {
  if (typeof onClick !== 'string') return onClick;
  return () => {
    navigate(onClick);
  };
}

function parseMenu(
  menus: Array<Menu>,
  navigate: NavigateFunction,
): Array<ParsedMenu> {
  return menus.map((menu) => ({
    ...menu,
    icon: parseIcon(menu.icon),
    onClick: parseOnClick(menu.onClick, navigate),
    subButton: menu.subButton && {
      ...menu.subButton,
      icon: parseIcon(menu.subButton.icon),
      onClick: parseOnClick(menu.subButton.onClick, navigate),
    },
    subMenu: parseMenu(menu.subMenu || [], navigate),
  }));
}

export function useSideMenu() {
  const [menu, setMenu] = useAtom(menuAtom);

  const navigate = useNavigate();

  return {
    setSideMenu(
      info: { title: string; desc: string; hideDesc: boolean },
      categories: Array<{ id: string; name: string }>,
      menus: Array<Menu>,
    ) {
      setMenu({
        info,
        categories,
        menus: parseMenu(menus, navigate),
      });
    },
    setSideMenuInfo(info: { title: string; desc: string; hideDesc: boolean }) {
      setMenu((prev) => {
        return {
          ...prev,
          info: info,
        };
      });
    },
    getSideMenu() {
      return menu;
    },
    getSideMenuInfo() {
      return menu.info;
    },
  };
}
