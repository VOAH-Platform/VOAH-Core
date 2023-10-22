import { styled } from '@/stitches.config';

export const ContextWrapper = styled('div', {
  display: 'flex',
  flexDirection: 'column',
  gap: '0.5rem',
  padding: '1rem',
  position: 'absolute',
  justifyContent: 'center',
  alignItems: 'flex-start',
  boxShadow: '0px 0px 8px 0px rgba(0, 0, 0, 0.15)',
  borderRadius: '0.75rem',
  whiteSpace: 'nowrap',
  background: '$gray0',
  zIndex: 100,

  variants: {
    isHidden: {
      true: {
        display: 'none',
      },
    },
  },
});

export const ContextCategoryWrapper = styled('div', {
  width: '100%',
  display: 'flex',
  flexDirection: 'column',
  alignItems: 'flex-start',
  justifyContent: 'center',
});

export const ContextCategoryName = styled('span', {
  marginBottom: '0.25rem',
  color: '$gray400',
  fontSize: '0.875rem',
  fontWeight: 500,
  letterSpacing: '-0.00875rem',
});

export const ContextCategoryButton = styled('button', {
  padding: '0.5rem',
  paddingRight: '1.5rem',
  width: '100%',
  display: 'flex',
  justifyContent: 'flex-start',
  alignItems: 'center',
  border: 'none',
  borderRadius: '0.5rem',
  background: '$gray0',
  fontSize: '1rem',
  color: '$gray600',
  fontWeight: 500,
  letterSpacing: '-0.01rem',
  cursor: 'pointer',
  transition: 'all 0.2s ease-in-out',

  '&:hover': {
    background: '$gray100',
  },
  '&:active': {
    color: '$gray800',
    fontWeight: 600,
  },

  '& > *': {
    stroke: '$gray600',
  },

  variants: {
    red: {
      true: {
        color: '$warning500',
        '& > *': {
          stroke: '$warning500',
        },
      },
    },
  },
});
