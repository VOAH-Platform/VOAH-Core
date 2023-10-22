import { styled } from '@/stitches.config';

export const VoahSideMenuWrapper = styled('div', {
  padding: '1rem',
  width: '16rem',
  height: '100%',
  display: 'flex',
  flexDirection: 'column',
  justifyContent: 'flex-start',
  alignItems: 'flex-start',
  gap: '0.75rem',
  borderRadius: '1.25rem',
  background: '$gray100',
});

export const VoahSideMenuTitleWrapper = styled('div', {
  width: '100%',
  display: 'flex',
  flexDirection: 'column',
  justifyContent: 'flex-start',
  alignItems: 'flex-start',
  gap: '0.25rem',
});

export const VoahSideMenuTitle = styled('span', {
  color: '$gray900',
  fontSize: '1.25rem',
  fontWeight: 700,
  letterSpacing: '-0.2px',
});

export const VoahSideMenuDesc = styled('span', {
  color: '$gray400',
  fontSize: '1rem',
  letterSpacing: '-0.16px',

  variants: {
    hidden: {
      true: {
        display: 'none',
      },
    },
  },
});

export const VoahSideMenuButton = styled('button', {
  padding: '10px 12px',
  width: '100%',
  display: 'flex',
  justifyContent: 'flex-start',
  alignItems: 'center',
  borderRadius: '0.5rem',
  background: '$gray100',
  border: 'none',
  gap: '0.5rem',
  cusor: 'pointer',

  '&:hover': {
    backgroundColor: '$gray200',
  },

  '&, & *': {
    cursor: 'pointer',
    stroke: '$gray500',
  },

  '& > span': {
    color: '$gray500',
    fontSize: '1.125rem',
    letterSpacing: '-0.18px',
  },
});
