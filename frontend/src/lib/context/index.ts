import { atom, useAtom } from 'jotai';

import { CustomContextProvider } from './provider';

export interface ContextCategory {
  nameHidden: boolean;
  id: string;
  name: string;
  buttons: Array<ContextButton>;
}

export interface ContextButton {
  icon: JSX.Element;
  name: string;
  onClick: (e: React.MouseEvent<HTMLButtonElement, MouseEvent>) => void;
  isRed?: boolean;
}

export const contextAtom = atom<{
  position: {
    x: number;
    y: number;
  };
  categories: Array<ContextCategory>;
}>({
  position: {
    x: 0,
    y: 0,
  },
  categories: [
    {
      nameHidden: false,
      id: 'test',
      name: '테스트 카테고리',
      buttons: [],
    },
  ],
});

export function useCustomContext() {
  const [, setContext] = useAtom(contextAtom);

  return {
    showContext: (x: number, y: number, categories: Array<ContextCategory>) => {
      setContext({
        position: {
          x,
          y,
        },
        categories,
      });
      document.dispatchEvent(new Event('voah__context_show'));
    },
    hideContext: () => {
      document.dispatchEvent(new Event('voah__context_hidden'));
    },
  };
}

export { CustomContextProvider };
