import { styled } from '@/stitches.config';

export const ContextWrapper = styled('div', {
  display: 'flex',
  flexDirection: 'column',
  padding: '1rem',
  position: 'absolute',
  justifyContent: 'center',
  alignItems: 'flex-start',
  boxShadow: '0px 0px 8px 0px rgba(0, 0, 0, 0.15)',
  borderRadius: '0.75rem',

  variants: {
    isHidden: {
      true: {
        display: 'none',
      },
    },
  },
});
