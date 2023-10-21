// import { atom } from 'jotai';
import { atom } from 'jotai';
import { atomWithStorage } from 'jotai/utils';

import { THEME_TOKEN } from './constant';

export const themeAtom = atomWithStorage('theme', {
  token: THEME_TOKEN.SYSTEM,
  isDark: false,
});

interface UserData {
  email: string;
  isLogin: boolean;
  id: string;
  accessToken: string;
  refreshToken: string;
}

export const userAtom = atomWithStorage<UserData>(
  'user',
  {
    email: '',
    isLogin: false,
    id: '',
    accessToken: '',
    refreshToken: '',
  },
  {
    getItem(key, initialValue) {
      const storedValue = localStorage.getItem(key);
      try {
        return JSON.parse(storedValue ?? '') as UserData;
      } catch {
        return initialValue;
      }
    },
    setItem(key, value) {
      localStorage.setItem(key, JSON.stringify(value));
    },
    removeItem(key) {
      localStorage.removeItem(key);
    },
    subscribe(key, callback, initialValue) {
      if (
        typeof window === 'undefined' ||
        typeof window.addEventListener === 'undefined'
      ) {
        return () => {};
      }
      window.addEventListener('storage', (e) => {
        if (e.storageArea === localStorage && e.key === key) {
          let newValue: UserData;
          try {
            newValue = JSON.parse(e.newValue ?? '') as UserData;
          } catch {
            newValue = initialValue;
          }
          callback(newValue);
        }
      });
      return () => {
        window.removeEventListener('storage', () => {});
      };
    },
  },
);

export const contextAtom = atom<{
  categories: Array<{
    nameHidden: boolean;
    id: string;
    name: string;
    buttons: Array<{
      icon: JSX.Element;
      name: string;
      onClick: (e: React.MouseEvent<HTMLButtonElement, MouseEvent>) => void;
    }>;
  }>;
}>({
  categories: [
    {
      nameHidden: false,
      id: 'test',
      name: '테스트 카테고리',
      buttons: [],
    },
  ],
});
