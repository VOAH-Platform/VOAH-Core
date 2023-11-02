import { atom } from 'jotai';
import { atomWithStorage } from 'jotai/utils';
import { HomeIcon } from 'lucide-react';

import { THEME_TOKEN } from './constant';

export const themeAtom = atomWithStorage('voah__theme', {
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
  'voah__user',
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

export const localDataAtom = atomWithStorage('voah__localData', {
  lastPath: '/app',
  useUndelineOnLink: false,
  isDevMode: false,
});

export const headerAtom = atom({
  isHidden: true,
  icon: <HomeIcon size={20} />,
  name: '메인',
});

export interface ModuleData {
  id: number;
  enabled: boolean;
  expose: boolean;
  version: string;
  name: string;
  description: string;
  'host-url': string;
  'permission-types': string;
  'permission-scopes': string;
  'created-at': string;
  'updated-at': string;
}

export const moduleAtom = atom({
  data: [] as Array<ModuleData>,
  indexMap: new Map<number, number>(),
});
