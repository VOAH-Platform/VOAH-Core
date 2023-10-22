import { styled } from '@/stitches.config';

export const VoahSideCategoryWrapper = styled('div', {
  padding: '0.75rem',
  height: '100%',
  borderRadius: '1.25rem',
  display: 'flex',
  flexDirection: 'column',
  justifyContent: 'flex-start',
  alignItems: 'center',
  gap: '1.25rem',
  background: '$gray100',
});

export const VoahSideCategoryTypeButton = styled('button', {
  border: 'none',
  borderRadius: '1rem',
  padding: '0.75rem',
  width: '3.25rem',
  height: '3.25rem',
  display: 'flex',
  justifyContent: 'center',
  alignItems: 'center',
  background: '$gray200',
  transition: 'background 0.2s ease-in-out',
  cursor: 'pointer',

  '&:hover': {
    background: '$gray300',
  },
});
