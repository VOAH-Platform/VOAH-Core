// import { atom } from 'jotai';
import { atomWithStorage } from 'jotai/utils';

import { THEME_TOKEN } from './constant';

export const themeAtom = atomWithStorage('theme', THEME_TOKEN.SYSTEM);
