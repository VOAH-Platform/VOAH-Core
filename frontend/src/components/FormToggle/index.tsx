import { useState } from 'react';

import { styled } from '@/stitches.config';

const ToggleLabel = styled('label', {
  position: 'relative',
  display: 'inline-block',
  width: '60px',
  height: '30px',
});

const ToggleInput = styled('input', {
  opacity: 0,
  width: 0,
  height: 0,

  '&:checked + span': {
    background: '$secondary400',
  },

  '&:checked + span:before': {
    transform: 'translateX(28px)',
  },
});

const ToggleSpan = styled('span', {
  position: 'absolute',
  cursor: 'pointer',
  top: 0,
  left: 0,
  right: 0,
  bottom: 0,
  background: '$gray200',
  transition: '.3s',
  borderRadius: '34px',

  '&:before': {
    position: 'absolute',
    content: '',
    height: '24px',
    width: '24px',
    left: '4px',
    bottom: '3px',
    background: '$gray0',
    borderRadius: '50%',
    transition: '.3s',
  },
});

export function FormToggle({
  toggled,
  onClick,
}: {
  toggled: boolean;
  onClick: (toggled: boolean) => void;
}) {
  const [isToggled, toggle] = useState(toggled);

  const callback = () => {
    toggle(!isToggled);
    onClick(!isToggled);
  };

  return (
    <ToggleLabel>
      <ToggleInput
        type="checkbox"
        defaultChecked={isToggled}
        onClick={callback}
      />
      <ToggleSpan />
    </ToggleLabel>
  );
}
