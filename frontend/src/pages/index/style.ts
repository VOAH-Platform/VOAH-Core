import { motion } from 'framer-motion';

import { styled } from '@/stitches.config';

export const IndexWrapper = styled(motion.div, {
  width: '100%',
  height: '100%',
  display: 'flex',
  flexDirection: 'column',
  justifyContent: 'center',
  alignItems: 'flex-start',
  gap: '12px',
  overflow: 'hidden',

  '& *': {
    userSelect: 'none',
    cursor: 'default',
  },
});

export const Container = styled(motion.div, {
  margin: '0 5rem',
  padding: '4rem 3.25rem',
  maxWidth: '520px',
  width: '100%',
  overflow: 'hidden',
  background: '$gray100',
  display: 'flex',
  flexDirection: 'column',
  gap: '2rem',
  borderRadius: '1.5rem',
  boxShadow: '$grade4',
  zIndex: 10,
});

export const ContainerHead = styled(motion.div, {
  display: 'flex',
  flexDirection: 'column',
  gap: '1rem',
});

export const ContainerTitle = styled('span', {
  color: '$gray900',
  fontSize: '2rem',
  fontWeight: 700,
  lineHeight: '140%',
  letterSpacing: '-0.02rem',
});

export const ContainerBody = styled(motion.div, {
  display: 'flex',
  flexDirection: 'column',
});

export const AnimWrapper = styled(motion.form, {
  display: 'flex',
  flexDirection: 'column',
  gap: '1.5rem',
});

export const FormError = styled(motion.span, {
  color: '$warning500',
  fontSize: '0.875rem',
  fontWeight: 500,
  lineHeight: '140%',
  letterSpacing: '-0.00875rem',
});

export const ButtonWrapper = styled(motion.div, {
  display: 'flex',
  gap: '0.75rem',
});

export const SmallAction = styled('span', {
  color: '$gray400',
  fontSize: '0.875rem',
  fontWeight: 500,
  lineHeight: '140%',
  letterSpacing: '-0.00875rem',
  cursor: 'pointer',
});

export const ThemeButton = styled('button', {
  position: 'fixed',
  bottom: '2rem',
  left: '2rem',
  padding: '0.5rem',
  background: '$transparent',
  border: 'none',
  textAlign: 'center',
  zIndex: 30,

  '&, & *': {
    cursor: 'pointer',
  },
});
