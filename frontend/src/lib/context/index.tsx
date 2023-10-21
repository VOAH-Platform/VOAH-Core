import { useAtom } from 'jotai';
import { useEffect, useRef, useState } from 'react';

import { contextAtom } from '@/atom';

import { ContextWrapper } from './style';

export function CustomContextProvider({
  children,
}: {
  children: React.ReactNode;
}) {
  const [Context] = useAtom(contextAtom);

  const [isHidden, setIsHidden] = useState(false);

  const contextRef = useRef<HTMLDivElement>(null);

  useEffect(() => {
    document.addEventListener('contextmenu', (e) => {
      e.preventDefault();
    });

    document.addEventListener('voah__context_show', (e) => {
      console.log(e);
      setIsHidden(false);
    });

    document.addEventListener('voah__context_hidden', (e) => {
      console.log(e);
      setIsHidden(true);
    });
  }, []);

  return (
    <>
      <ContextWrapper ref={contextRef} isHidden={isHidden}>
        {Context.categories.map((category, key) => (
          <div key={key}>
            <span>{category.name}</span>
            {category.buttons.map((button, key2) => (
              <button key={key2} onClick={button.onClick}>
                {button.name}
              </button>
            ))}
          </div>
        ))}
      </ContextWrapper>
      {children}
    </>
  );
}
