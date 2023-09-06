import { motion } from 'framer-motion';

import { styled } from '@/stitches.config';

export const SelectWrapper = styled(motion.div, {
  display: 'flex',
  flexDirection: 'column',
  gap: '0.75rem',

  '& *': {
    userSelect: 'none',
  },
});

export const SelectLable = styled('label', {
  color: '$gray600',
  fontSize: '1.125rem',
  fontWeight: 600,
  lineHeight: '140%',
  letterSpacing: '-0.01125rem',
});

export const OptionWrapper = styled('div', {
  padding: '0.75rem 1rem',
  borderRadius: '0.63rem',
  display: 'flex',
  alignItems: 'center',
  justifyContent: 'space-between',

  '&:hover': {
    background: '$secondary200',
    '& *': {
      color: '$gray900',
    },
  },
});

export const OptionText = styled('span', {
  color: '$gray700',
  fontSize: '1.125rem',
  fontWeight: 600,
  lineHeight: '140%',
  letterSpacing: '-0.01125rem',
});

export const OptionPublic = styled('div', {
  display: 'flex',
  alignItems: 'center',
  gap: '0.25rem',
  '& *': {
    color: '$gray500',
  },
});
