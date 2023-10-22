import { styled } from '@/stitches.config';

export const VoahSettingsProfileWrapper = styled('div', {
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

export const VoahSettingsProfileHead = styled('div', {
  width: '100%',
  display: 'flex',
  justifyContent: 'space-between',
  alignItems: 'center',
});

export const VoahSettingsProfileUserInfoWrapper = styled('div', {
  display: 'flex',
  justifyContent: 'flex-start',
  alignItems: 'center',
  gap: '1.5rem',
});

export const VoahSettingsProfileUserInfoImage = styled('div', {
  width: '8rem',
  height: '8rem',
  borderRadius: '50%',
  overflow: 'hidden',
  background: '$gray300',
});

export const VoahSettingsProfileUserInfoTextWrapper = styled('div', {
  display: 'flex',
  flexDirection: 'column',
  justifyContent: 'center',
  alignItems: 'flex-start',
});

export const VoahSettingsProfileUserInfoDisplayName = styled('span', {
  color: '$gray700',
  fontSize: '2rem',
  fontWeight: 600,
  letterSpacing: '-0.32px',
});

export const VoahSettingsProfileUserInfoUsername = styled('span', {
  color: '$gray400',
  fontSize: '1.5rem',
  letterSpacing: '-0.24px',
});

export const VoahSettingsProfileEditButton = styled('button', {
  padding: '0.5rem 0.75rem',
  display: 'flex',
  justifyContent: 'center',
  alignItems: 'center',
  gap: '0.25rem',
  border: 'none',
  borderRadius: '0.75rem',
  background: '$gray100',
  transition: 'all 0.2s ease-in-out',
  cursor: 'pointer',

  '&:hover': {
    background: '$gray200',
  },
});

export const VoahSettingsProfileOtherWrapper = styled('div', {
  display: 'flex',
  flexDirection: 'column',
  justifyContent: 'center',
  alignItems: 'flex-start',
  gap: '0.5rem',
});

export const VoahSettingsProfileOtherTitleWrapper = styled('div', {
  display: 'flex',
  justifyContent: 'flex-start',
  alignItems: 'center',
  gap: '0.5rem',

  '& *': {
    stroke: '$gray400',
  },
});

export const VoahSettingsProfileOtherTitle = styled('span', {
  color: '$gray400',
  fontSize: '1.25rem',
  letterSpacing: '-0.2px',
});

export const VoahSettingsProfileOtherValue = styled('span', {
  color: '$gray700',
  fontSize: '1.75rem',
  fontWeight: 600,
  letterSpacing: '-0.28px',
});
