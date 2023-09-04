import { Input, InputLable, InputWrapper } from './style';

export function FormInput({
  id,
  label,
  onChange,
  ...props
}: {
  id: string;
  label: string;
  onChange?: (e: React.ChangeEvent<HTMLInputElement>) => void;
  [x: string]: unknown;
}) {
  return (
    <InputWrapper layout>
      <InputLable htmlFor={id}>{label}</InputLable>
      <Input id={id} onChange={onChange} {...props} />
    </InputWrapper>
  );
}
