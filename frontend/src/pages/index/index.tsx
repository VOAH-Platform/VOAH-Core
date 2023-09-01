import { useAtom } from 'jotai';

import { themeAtom } from '@/atom';
import { ExampleButton } from '@/components/ExampleButton';
import { THEME_TOKEN } from '@/constant';

import { IndexWrapper } from './style';

export function IndexPage() {
  const [, setTheme] = useAtom(themeAtom);

  return (
    <IndexWrapper>
      VOAH TEMPLATE
      <ExampleButton onClick={() => setTheme(THEME_TOKEN.LIGHT)}>
        LIGHT
      </ExampleButton>
      <ExampleButton onClick={() => setTheme(THEME_TOKEN.DARK)}>
        DARK
      </ExampleButton>
      <ExampleButton onClick={() => setTheme(THEME_TOKEN.SYSTEM)}>
        SYSTEM
      </ExampleButton>
    </IndexWrapper>
  );
}
