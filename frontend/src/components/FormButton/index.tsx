import { styled } from '@/stitches.config';

const Button = styled('button', {
  padding: '1rem 3rem',
  background: '$gray100',
  fontSize: '1.125rem',
  fontWeight: 700,
  lineHeight: '140%',
  letterSpacing: '-0.01125rem',
  borderRadius: '1rem',
  color: '$secondary400',
  border: '2px solid $secondary400',
  cursor: 'pointer',
  '&:hover': {
    background: '$gray200',
  },
  '&[disabled]': {
    border: '2px solid $gray400',
    color: '$gray400',
    cursor: 'not-allowed',

    '&:hover': {
      background: '$gray100',
    },
  },

  variants: {
    filled: {
      true: {
        background: '$secondary400',
        color: '$gray0',
        '&:hover': {
          background: '$secondary500',
        },
        '&[disabled]': {
          border: '2px solid $gray400',
          background: '$gray400',
          color: '$gray100',
          cursor: 'not-allowed',
          '&:hover': {
            background: '$gray400',
          },
        },
      },
    },
  },
});

export function FormButton({
  children,
  onClick,
  ...props
}: {
  children: React.ReactNode;
  onClick?: (e: React.MouseEvent<HTMLButtonElement>) => void;
  [x: string]: unknown;
}) {
  return (
    <Button onClick={onClick} {...props}>
      {children}
    </Button>
  );
}
