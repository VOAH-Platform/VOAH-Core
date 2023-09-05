import { motion } from 'framer-motion';

import { styled } from '@/stitches.config';

export const InputWrapper = styled(motion.div, {
  display: 'flex',
  flexDirection: 'column',
  gap: '0.75rem',
});

export const InputLable = styled('label', {
  color: '$gray600',
  fontSize: '1.125rem',
  fontWeight: 600,
  lineHeight: '140%',
  letterSpacing: '-0.01125rem',
});

export const Input = styled('input', {
  padding: '1.25rem 1.5rem',
  color: '$gray900',
  fontSize: '1.125rem',
  fontWeight: 500,
  lineHeight: '140%',
  letterSpacing: '-0.01125rem',
  borderRadius: '1rem',
  border: '2px solid $gray200',
  backgroundColor: '$gray0',
  cursor: 'text',

  '&:focus': {
    outline: 'none',
    border: '2px solid $secondary400',
  },
});
