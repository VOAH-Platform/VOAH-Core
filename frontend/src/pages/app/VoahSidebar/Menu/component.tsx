import { useAtom } from 'jotai';
import { useEffect, useState } from 'react';

import {
  VoahSideMenuButton,
  VoahSideMenuDesc,
  VoahSideMenuTitle,
  VoahSideMenuTitleWrapper,
  VoahSideMenuWrapper,
} from './style';

import { ParsedMenu, menuAtom } from '.';

function parseMenuToTree(
  categories: Array<{ id: string; name: string }>,
  menus: Array<ParsedMenu>,
): {
  solo: Array<ParsedMenu>;
  tree: Array<{
    id: string;
    name: string;
    menus: Array<ParsedMenu>;
  }>;
} {
  const solo: Array<ParsedMenu> = [];
  const tree: Array<{
    id: string;
    name: string;
    menus: Array<ParsedMenu>;
  }> = categories.map((category) => ({
    id: category.id,
    name: category.name,
    menus: [] as Array<ParsedMenu>,
  }));

  menus.forEach((menu) => {
    if (menu.categoryId) {
      const target = tree.find((item) => item.id === menu.categoryId);
      if (target) target.menus.push(menu);
    } else {
      solo.push(menu);
    }
  });

  return {
    solo,
    tree,
  };
}

export function VoahSideMenu() {
  const [menuData] = useAtom(menuAtom);

  const [menu, setMenu] = useState<{
    solo: Array<ParsedMenu>;
    tree: Array<{
      id: string;
      name: string;
      menus: Array<ParsedMenu>;
    }>;
  }>(parseMenuToTree(menuData.categories, menuData.menus));

  useEffect(() => {
    setMenu(parseMenuToTree(menuData.categories, menuData.menus));
  }, [menuData]);

  return (
    <VoahSideMenuWrapper>
      <VoahSideMenuTitleWrapper>
        <VoahSideMenuTitle>{menuData.info.title}</VoahSideMenuTitle>
        <VoahSideMenuDesc hidden={menuData.info.hideDesc}>
          {menuData.info.desc}
        </VoahSideMenuDesc>
      </VoahSideMenuTitleWrapper>
      {menu.solo.map((menu, index) => (
        <VoahSideMenuButton key={index} onClick={menu.onClick}>
          {menu.icon}
          <span>{menu.name}</span>
        </VoahSideMenuButton>
      ))}
      {menu.tree.map((tree, index) => (
        <div key={index}>{tree.name}</div>
      ))}
    </VoahSideMenuWrapper>
  );
}
