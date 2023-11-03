import { useAtom } from 'jotai';
import { CheckIcon, MonitorIcon, PaletteIcon, Paperclip } from 'lucide-react';

import { localDataAtom, themeAtom } from '@/atom';
import { FormToggle } from '@/components/FormToggle';
import { THEME_TOKEN } from '@/constant';

import { useSettingsInit } from '../logic';

import {
  VoahSettingsAccessibilityMenuLink,
  VoahSettingsAccessibilityMenuRow,
  VoahSettingsAccessibilityMenuTitle,
  VoahSettingsAccessibilityMenuWrapper,
  VoahSettingsAccessibilityWrapper,
  VoahSettingsThemeButton,
  VoahSettingsThemeWrapper,
} from './style';

export function AccessibilityPage() {
  const [theme, setTheme] = useAtom(themeAtom);
  const [localData, setLocalData] = useAtom(localDataAtom);

  useSettingsInit();

  return (
    <VoahSettingsAccessibilityWrapper>
      <VoahSettingsAccessibilityMenuWrapper>
        <VoahSettingsAccessibilityMenuTitle>
          <PaletteIcon /> 테마
        </VoahSettingsAccessibilityMenuTitle>
        <VoahSettingsAccessibilityMenuRow>
          <VoahSettingsThemeWrapper>
            <VoahSettingsThemeButton
              type="light"
              onClick={() => {
                setTheme({
                  token: THEME_TOKEN.LIGHT,
                  isDark: false,
                });
              }}>
              {theme.token == THEME_TOKEN.LIGHT && <CheckIcon />}
            </VoahSettingsThemeButton>
            <span>라이트</span>
          </VoahSettingsThemeWrapper>
          <VoahSettingsThemeWrapper>
            <VoahSettingsThemeButton
              type="dark"
              onClick={() => {
                setTheme({
                  token: THEME_TOKEN.DARK,
                  isDark: true,
                });
              }}>
              {theme.token == THEME_TOKEN.DARK && <CheckIcon />}
            </VoahSettingsThemeButton>
            <span>다크</span>
          </VoahSettingsThemeWrapper>
          <VoahSettingsThemeWrapper>
            <VoahSettingsThemeButton
              type="system"
              onClick={() => {
                setTheme({
                  token: THEME_TOKEN.SYSTEM,
                  isDark: false,
                });
              }}>
              {theme.token == THEME_TOKEN.SYSTEM && <CheckIcon />}
            </VoahSettingsThemeButton>
            <span>시스템</span>
          </VoahSettingsThemeWrapper>
        </VoahSettingsAccessibilityMenuRow>
      </VoahSettingsAccessibilityMenuWrapper>
      <hr />
      <VoahSettingsAccessibilityMenuWrapper>
        <VoahSettingsAccessibilityMenuTitle>
          <Paperclip /> 링크에 밑줄 표시
        </VoahSettingsAccessibilityMenuTitle>
        <VoahSettingsAccessibilityMenuRow>
          <VoahSettingsAccessibilityMenuLink
            underline={localData.useUndelineOnLink}>
            https://implude.kr
          </VoahSettingsAccessibilityMenuLink>
          <FormToggle
            toggled={localData.useUndelineOnLink}
            onClick={() => {
              setLocalData((prev) => ({
                ...prev,
                useUndelineOnLink: !prev.useUndelineOnLink,
              }));
            }}
          />
        </VoahSettingsAccessibilityMenuRow>
      </VoahSettingsAccessibilityMenuWrapper>
      <hr />
      <VoahSettingsAccessibilityMenuWrapper>
        <VoahSettingsAccessibilityMenuTitle>
          <MonitorIcon /> 개발자 모드
        </VoahSettingsAccessibilityMenuTitle>
        <VoahSettingsAccessibilityMenuRow>
          <FormToggle
            toggled={localData.isDevMode}
            onClick={(state) => {
              setLocalData((prev) => ({
                ...prev,
                isDevMode: state,
              }));
            }}
          />
        </VoahSettingsAccessibilityMenuRow>
      </VoahSettingsAccessibilityMenuWrapper>
    </VoahSettingsAccessibilityWrapper>
  );
}
