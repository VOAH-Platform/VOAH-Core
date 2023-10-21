import { styled } from '@/stitches.config';

export const VoahFrameWrapper = styled('div', {
  padding: '1rem',
  display: 'flex',
  flexDirection: 'column',
  height: '100%',
  width: '100%',
  flex: '1',
  gap: '1rem',
});

export const AddressBarWrapper = styled('div', {
  display: 'flex',
  alignItems: 'center',
  gap: '0.5rem',
});

export const AddressBar = styled('input', {
  flex: '1',
  height: '2.5rem',
  borderRadius: '0.75rem',
  border: '2px solid $gray300',
  padding: '0.5rem',
  outline: 'none',
  fontSize: '1rem',
  color: '$gray700',

  '&:focus': {
    border: '1px solid $gray600',
  },
});

export const AddressBarBtn = styled('button', {
  height: '2.5rem',
  border: 'none',
  borderRadius: '0.75rem',
  background: '$gray400',
  padding: '0.5rem',
  outline: 'none',
  fontSize: '1rem',
  color: '$gray0',
  transition: 'background 0.2s ease-in-out',
  cursor: 'pointer',

  '&:hover': {
    background: '$gray500',
  },
});
