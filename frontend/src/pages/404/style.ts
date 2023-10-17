import { styled } from '@stitches/react';

export const NotFoundWrapper = styled('div', {
  height: '100%',
  display: 'flex',
  flexDirection: 'column',
  justifyContent: 'center',
  alignItems: 'center',
});

export const NotFoundSubtitle = styled('span', {
  margin: '0.5rem 0 1.5rem 0',
  fontSize: '1.5rem',
  fontWeight: '$normal',
  lineHeight: '140%',
  letterSpacing: '-0.015rem',
  color: '$gray400',
});

export const NotFoundButton = styled('a', {
  display: 'flex',
  padding: '0.5rem 0.75rem',
  alignItems: 'center',
  gap: '0.25rem',
  borderRadius: '0.75rem',
  background: '$gray100',
  color: '$gray600',
  fontSize: '1rem',
  fontWeight: '600',
  lineHeight: '140%',
  letterSpacing: '-0.01rem',
  cursor: 'pointer',
  textDecoration: 'none',
});
