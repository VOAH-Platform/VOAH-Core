import { atom, useAtom } from 'jotai';
import { KeyRoundIcon } from 'lucide-react';

import { LucideCustom } from '@/lib/LucideCustom';

export interface Menu {
  icon: JSX.Element | string;
  name: string;
  onClick: (e: React.MouseEvent<HTMLButtonElement>) => void;
  mentioned?: number;
  isFocused?: boolean;
  categoryId?: string;
  subMenu?: Array<Menu>;
}

export interface ParsedMenu extends Menu {
  icon: JSX.Element;
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

function parseMenu(menus: Array<Menu>): Array<ParsedMenu> {
  return menus.map((menu) => ({
    ...menu,
    icon: parseIcon(menu.icon),
    subMenu: parseMenu(menu.subMenu || []),
  }));
}

export function useSideMenu() {
  const [, setMenu] = useAtom(menuAtom);

  return {
    setSideMenu(
      info: { title: string; desc: string; hideDesc: boolean },
      categories: Array<{ id: string; name: string }>,
      menus: Array<Menu>,
    ) {
      setMenu({
        info,
        categories,
        menus: parseMenu(menus),
      });
    },
  };
}
