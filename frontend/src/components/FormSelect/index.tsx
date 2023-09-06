import { Lock, Users2 } from 'lucide-react';
import { useState } from 'react';
import Select, { OptionProps } from 'react-select';

import { css } from '@/stitches.config';

import {
  OptionText,
  OptionWrapper,
  OptionPublic,
  SelectLable,
  SelectWrapper,
} from './style';

function CustomFormOption({
  isSelected,
  selectOption,
  label,
  data,
}: OptionProps<{
  value: string;
  label: string;
  private?: boolean;
}>) {
  return (
    <OptionWrapper
      onClick={() => {
        selectOption(data);
      }}
      className={
        isSelected
          ? css({
              background: '$secondary400 !important',
              '& *': {
                color: '$gray0 !important',
              },
            })().className
          : ''
      }>
      <OptionText>{label}</OptionText>
      <OptionPublic>
        {data.private ? (
          <>
            <Lock />
            <span>초대 받음</span>
          </>
        ) : (
          <>
            <Users2 />
            <span>공개</span>
          </>
        )}
      </OptionPublic>
    </OptionWrapper>
  );
}

export function FormSelect({
  id,
  label,
  items,
  onSelect,
  ...props
}: {
  id: string;
  label: string;
  items: {
    value: string;
    label: string;
    private?: boolean;
  }[];
  onSelect?: (val: { value: string; label: string }) => void;
  [x: string]: unknown;
}) {
  const [selectedOption, setSelectedOption] = useState<{
    value: string;
    label: string;
  } | null>(null);

  return (
    <SelectWrapper layout>
      <SelectLable htmlFor={id}>{label}</SelectLable>
      <Select
        defaultValue={selectedOption}
        onChange={(val) => {
          setSelectedOption(val);
          onSelect?.(val!);
        }}
        isMulti={false}
        options={items}
        components={{
          Option: CustomFormOption,
        }}
        classNames={{
          control: (state) =>
            !state.isFocused
              ? css({
                  padding: '0.9rem 1rem !important',
                  color: '$gray900 !important',
                  fontSize: '1.125rem !important',
                  lineHeight: '140% !important',
                  letterSpacing: '-0.01125rem !important',
                  borderRadius: '1rem !important',
                  border: '2px solid $gray200 !important',
                  backgroundColor: '$gray0 !important',
                })().className
              : css({
                  padding: '0.9rem 1rem !important',
                  color: '$gray900 !important',
                  fontSize: '1.125rem !important',
                  lineHeight: '140% !important',
                  letterSpacing: '-0.01125rem !important',
                  borderRadius: '1rem !important',
                  border: '2px solid $secondary400 !important',
                  backgroundColor: '$gray0 !important',
                  boxShadow: 'none !important',
                })().className,
          menu: () =>
            css({
              marginTop: '1rem !important',
              background: '$gray0 !important',
              borderRadius: '1rem !important',
              border: '2px solid $gray200 !important',
            })().className,
          menuList: () =>
            css({
              padding: '0.75rem !important',
              background: '$gray0 !important',
              borderRadius: '1rem !important',
            })().className,
          singleValue: () =>
            css({
              color: '$gray900 !important',
            })().className,
        }}
        {...props}
      />
    </SelectWrapper>
  );
}
