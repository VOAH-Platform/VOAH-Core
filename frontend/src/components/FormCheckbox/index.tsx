import { Check } from 'lucide-react';
import { useState } from 'react';

import { styled } from '@/stitches.config';

const CheckboxWrapper = styled('div', {
  display: 'flex',
  alignItems: 'center',
  gap: '0.5rem',
  cursor: 'pointer',
  userSelect: 'none',
  '&[disabled]': {
    opacity: 0.5,
    cursor: 'not-allowed',
  },
});

const CheckboxFake = styled('div', {
  margin: '0.125rem',
  width: '1.75rem',
  height: '1.75rem',
  borderRadius: '0.5rem',
  border: '2px solid $gray700',
  backgroundColor: '$gray0',
  display: 'flex',
  alignItems: 'center',
  justifyContent: 'center',

  '& *': {
    color: '$gray0',
  },

  variants: {
    checked: {
      true: {
        backgroundColor: '$gray700',
        color: '$gray0',
      },
    },
  },
});

const CheckboxLabel = styled('label', {
  color: '$gray700',
  fontSize: '1.25rem',
  fontWeight: 500,
  lineHeight: '140%',
  letterSpacing: '-0.0125rem',
  userSelect: 'none',
  cursor: 'pointer',
});

export function FormCheckbox({
  id,
  text,
  onChange,
  ...props
}: {
  id: string;
  text: string;
  onChange?: (e: React.ChangeEvent<HTMLInputElement>) => void;
}) {
  const [checked, setChecked] = useState(false);

  return (
    <CheckboxWrapper>
      <input
        id={id}
        checked={checked}
        onChange={(e) => {
          setChecked(e.target.checked);
          onChange?.(e);
        }}
        type="checkbox"
        hidden
        {...props}
      />
      <CheckboxFake
        onClick={() => {
          setChecked(!checked);
          onChange?.({
            target: {
              checked: !checked,
            },
          } as React.ChangeEvent<HTMLInputElement>);
        }}
        checked={checked}>
        <Check />
      </CheckboxFake>
      <CheckboxLabel htmlFor={id}>{text}</CheckboxLabel>
    </CheckboxWrapper>
  );
}
