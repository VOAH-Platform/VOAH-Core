import { CheckCircle2, Info, XCircle, XIcon } from 'lucide-react';
import { AlertTemplateProps } from 'react-alert';

import { AlertButton, AlertIcon, AlertText, AlertWrapper } from './style';

export function AlertTemplate({
  style,
  options,
  message,
  close,
}: AlertTemplateProps) {
  return (
    <AlertWrapper style={style} type={options.type}>
      <AlertIcon>
        {options.type == 'info' && <Info />}
        {options.type == 'error' && <XCircle />}
        {options.type == 'success' && <CheckCircle2 />}
      </AlertIcon>
      <AlertText>{message}</AlertText>
      <AlertButton onClick={close}>
        <XIcon />
      </AlertButton>
    </AlertWrapper>
  );
}
