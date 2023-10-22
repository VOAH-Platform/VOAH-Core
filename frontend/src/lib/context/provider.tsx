import { useAtom } from 'jotai';
import { useEffect, useRef, useState } from 'react';

import {
  ContextCategoryButton,
  ContextCategoryName,
  ContextCategoryWrapper,
  ContextWrapper,
} from './style';

import { contextAtom } from '.';

export function CustomContextProvider({
  children,
}: {
  children: React.ReactNode;
}) {
  const [Context] = useAtom(contextAtom);

  const [isHidden, setIsHidden] = useState(true);

  const contextRef = useRef<HTMLDivElement>(null);

  const [realPosition, setRealPosition] = useState({
    x: 0,
    y: 0,
  });

  useEffect(() => {
    document.addEventListener('contextmenu', (e) => {
      e.preventDefault();
    });

    document.addEventListener('voah__context_show', () => {
      setIsHidden(false);

      // dispatch hidden event on click outside of context menu
      document.removeEventListener('mousedown', () => {});
      document.addEventListener(
        'mousedown',
        (e) => {
          if (contextRef.current) {
            if (!contextRef.current.contains(e.target as Node)) {
              document.dispatchEvent(new Event('voah__context_hidden'));
            }
          }
        },
        { once: true },
      );
    });

    document.addEventListener('voah__context_hidden', () => {
      setIsHidden(true);
    });
  }, []);

  useEffect(() => {
    setRealPosition({
      x: Context.position.x,
      y: Context.position.y,
    });

    // calculate context menu position by contextatom's position and context menu size
    setTimeout(() => {
      if (contextRef.current) {
        const contextWidth = contextRef.current.offsetWidth;
        const contextHeight = contextRef.current.offsetHeight;

        const windowWidth = window.innerWidth;
        const windowHeight = window.innerHeight;

        const { x, y } = Context.position;

        const realX = x + contextWidth > windowWidth ? x - contextWidth : x;
        const realY = y + contextHeight > windowHeight ? y - contextHeight : y;

        setRealPosition({
          x: realX,
          y: realY,
        });
      }
    }, 2);
  }, [Context, contextRef.current]);

  useEffect(() => {
    setIsHidden(true);
    document.dispatchEvent(new Event('voah__context_hidden'));
  }, [window.location.href]);

  return (
    <>
      <ContextWrapper
        style={{
          left: realPosition.x,
          top: realPosition.y,
        }}
        ref={contextRef}
        isHidden={isHidden}>
        {Context.categories.map((category, key) => (
          <ContextCategoryWrapper key={key}>
            <ContextCategoryName>{category.name}</ContextCategoryName>
            {category.buttons.map((button, key2) => (
              <ContextCategoryButton
                key={key2}
                onClick={button.onClick}
                red={button.isRed}>
                {button.icon && <>{button.icon}&nbsp;&nbsp;</>}
                {button.name}
              </ContextCategoryButton>
            ))}
          </ContextCategoryWrapper>
        ))}
      </ContextWrapper>
      {children}
    </>
  );
}
