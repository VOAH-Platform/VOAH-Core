import { Menu, useSideMenu } from '../VoahSidebar/Menu';

export function useVoahMessageFunc(port1: MessagePort) {
  return {
    frame: {
      initDone: (url: string) => {
        console.log(`${url} is loaded!`);
      },
    },
    user: {
      getToken: (accessToken: string) => {
        port1.postMessage({
          type: 'VOAH__USER_GET_TOKEN_DONE',
          data: accessToken,
        });
      },
    },
    sidebar: {
      setSidebarInfo: (data: {
        title: string;
        desc: string;
        hideDesc: boolean;
      }) => {
        const sideMenu = useSideMenu();
        sideMenu.setSideMenuInfo(data);
      },
      setSidebarMenu: (data: {
        categories: Array<{ id: string; name: string }>;
        menus: Array<Menu>;
      }) => {
        const sideMenu = useSideMenu();
        const info = sideMenu.getSideMenuInfo();
        sideMenu.setSideMenu(info, data.categories, data.menus);
      },
    },
  };
}
