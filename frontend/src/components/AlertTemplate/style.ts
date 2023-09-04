import { styled } from '@/stitches.config';

export const AlertWrapper = styled('div', {
  padding: '0.75rem 1rem',
  maxWidth: '450px',
  display: 'flex',
  justifyContent: 'center',
  alignItems: 'center',
  gap: '1rem',
  background: '$gray400',
  borderRadius: '0.5rem',
  boxShadow: '$grade2',
  '& *': {
    color: '$gray0',
  },

  variants: {
    type: {
      info: {},
      success: {
        background: '$accept400',
        '& *': {
          color: '$gray0',
        },
      },
      error: {
        background: '$warning400',
        '& *': {
          color: '$gray0',
        },
      },
    },
  },
});

export const AlertIcon = styled('div', {
  width: '24px',
  height: '24px',
  display: 'flex',
  justifyContent: 'center',
  alignItems: 'center',
  '& svg': {
    width: '100%',
    height: '100%',
  },
});

export const AlertText = styled('span', {
  flex: 1,
  fontSize: '1rem',
  fontWeight: 700,
  lineHeight: '140%',
  letterSpacing: '-0.00875rem',
  textAlign: 'left',
});

export const AlertButton = styled('button', {
  padding: '0',
  display: 'flex',
  justifyContent: 'center',
  alignItems: 'center',
  border: 'none',
  background: 'none',
  '&, & *': {
    cursor: 'pointer',
  },
});
