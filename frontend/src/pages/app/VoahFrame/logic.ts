import { apiClient } from '@/apiClient';
import { UserData } from '@/atom';

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
      getUser: (user: UserData) => {
        port1.postMessage({
          type: 'VOAH__USER_GET_USER_DONE',
          data: user,
        });
      },
      getProfile: async (accessToken: string, id: string) => {
        const response = await apiClient.get(`/api/profile?user-id=${id}`, {
          headers: {
            Authorization: `Bearer ${accessToken}`,
          },
        });
        const data = {
          ...response.data,
          accessToken,
        } as {
          user: {
            'user-id': string;
            username: string;
            email: string;
          };
          accessToken: string;
        };
        console.log('profile data', data);
        port1.postMessage({
          type: 'VOAH__USER_GET_PROFILE_DONE',
          data: data,
        });
      },
    },
    sidebar: {
      useSidebarInfo: (data: {
        title: string;
        desc: string;
        hideDesc: boolean;
      }) => {
        const sideMenu = useSideMenu();
        sideMenu.setSideMenuInfo(data);
      },
      useSidebarMenu: (data: {
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
