import { styled } from '@/stitches.config';

export const VoahSettingsAccessibilityWrapper = styled('div', {
  padding: '1.25rem 2rem 1.25rem 1.25rem',
  width: '100%',
  height: '100%',
  display: 'flex',
  flexDirection: 'column',
  justifyContent: 'flex-start',
  alignItems: 'flex-start',
  gap: '1.75rem',

  '& hr': {
    margin: '0',
    width: '100%',
    border: 'none',
    borderBottom: '1px solid $gray300',
  },
});

export const VoahSettingsAccessibilityMenuWrapper = styled('div', {
  width: '100%',
  display: 'flex',
  flexDirection: 'column',
  justifyContent: 'flex-start',
  alignItems: 'flex-start',
  gap: '1rem',
});

export const VoahSettingsAccessibilityMenuTitle = styled('span', {
  display: 'flex',
  justifyContent: 'flex-start',
  alignItems: 'center',
  gap: '0.5rem',
  color: '$gray400',
  fontSize: '1.25rem',
  letterSpacing: '-0.2px',

  '& svg': {
    width: '20px',
    height: '20px',
    stroke: '$gray400',
  },
});

export const VoahSettingsAccessibilityMenuRow = styled('div', {
  width: '100%',
  display: 'flex',
  justifyContent: 'flex-start',
  alignItems: 'center',
  gap: '1rem',
});

export const VoahSettingsThemeWrapper = styled('div', {
  display: 'flex',
  flexDirection: 'column',
  gap: '0.5rem',
  justifyContent: 'flex-start',
  alignItems: 'center',
});

export const VoahSettingsThemeButton = styled('button', {
  border: 'none',
  borderRadius: '50%',
  cursor: 'pointer',
  width: '3rem',
  height: '3rem',

  variants: {
    type: {
      light: {
        background: '$gray200',
      },
      dark: {
        background: '$gray700',
        '& svg': {
          stroke: '$gray0',
        },
      },
      system: {
        background: '$gray0',
        border: '2px solid $gray400',
        '& svg': {
          stroke: '$gray400',
        },
      },
    },
  },
});

export const VoahSettingsAccessibilityMenuLink = styled('span', {
  color: '$gray700',
  fontSize: '1.5rem',
  letterSpacing: '-0.24px',
  variants: {
    underline: {
      true: {
        textDecoration: 'underline',
      },
      false: {
        textDecoration: 'none',
      },
    },
  },
});
