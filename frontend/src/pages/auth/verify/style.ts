import { styled } from '@/stitches.config';

export const RegisterWrapper = styled('div', {
  width: '100%',
  padding: '3rem',
  display: 'flex',
  flexDirection: 'column',
});

export const RegisterHeader = styled('div', {
  width: '100%',
  display: 'flex',
  justifyContent: 'space-between',
});

export const ThemeButton = styled('button', {
  background: '$transparent',
  border: 'none',
  textAlign: 'center',

  '&, & *': {
    cursor: 'pointer',
  },
});

export const RegisterBody = styled('div', {
  marginTop: '2rem',
  padding: '0 2rem',
  width: '100%',
  maxWidth: '650px',
  display: 'flex',
  flexDirection: 'column',
});

export const RegisterTitle = styled('h1', {
  color: '$gray900',
  fontSize: '2.25rem',
  fontWeight: 700,
  lineHeight: '140%',
  letterSpacing: '-0.0225rem',
});

export const RegisterForm = styled('form', {
  marginTop: '3rem',
  display: 'flex',
  flexDirection: 'column',
  gap: '2rem',
});

export const TosWrapper = styled('div', {
  display: 'flex',
  flexDirection: 'column',
  gap: '0.75rem',
});

export const TosTitle = styled('span', {
  color: '$gray600',
  fontSize: '1.125rem',
  fontWeight: 600,
  lineHeight: '140%',
  letterSpacing: '-0.01125rem',
});

export const TosLinkWrapper = styled('div', {
  display: 'flex',
  gap: '0.5rem',
  alignItems: 'center',
  color: '$gray700',
  fontSize: '1.25rem',
  fontWeight: 600,
  lineHeight: '140%',
  letterSpacing: '-0.0125rem',
  textDecoration: 'underline',
  '& svg': {
    width: '2rem',
    height: '2rem',
  },
});
