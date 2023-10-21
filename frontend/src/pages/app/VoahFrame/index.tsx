import { useAtom } from 'jotai';
import { useEffect, useRef, useState } from 'react';

import { userAtom } from '@/atom';

import {
  AddressBar,
  AddressBarBtn,
  AddressBarWrapper,
  VoahFrameWrapper,
} from './style';

export function VoahFrame() {
  const [user] = useAtom(userAtom);

  // const [devMode, setDevMode] = useState(true);
  const [url, setUrl] = useState('');

  const frameRef = useRef<HTMLIFrameElement>(null);

  const { port1, port2 } = new MessageChannel();

  const [inputVal, setInputVal] = useState('');

  useEffect(() => {
    frameRef.current?.contentWindow?.postMessage('VOAH__FRAME_INIT', '*', [
      port2,
    ]);

    port1.onmessage = (
      e: MessageEvent<{
        type: string;
        data: unknown;
      }>,
    ) => {
      console.log(e.data);
      switch (e.data.type) {
        case 'VOAH__FRAME_INIT_DONE':
          console.log(`${url} is loaded!`);
          break;
        case 'VOAH__USER_GET_TOKEN':
          port1.postMessage({
            type: 'VOAH__USER_GET_TOKEN_DONE',
            data: user.accessToken,
          });
          break;
      }
    };
  }, [url]);

  return (
    <VoahFrameWrapper>
      <AddressBarWrapper>
        <AddressBar
          value={inputVal}
          onChange={(e) => {
            setInputVal(e.target.value);
          }}
          placeholder="모듈 URL을 입력하세요"
        />
        <AddressBarBtn
          onClick={() => {
            setUrl(inputVal);
          }}>
          이동하기
        </AddressBarBtn>
      </AddressBarWrapper>
      <iframe
        style={{
          width: '100%',
          height: '100%',
          border: 'none',
        }}
        ref={frameRef}
        title="module"
        src={url}></iframe>
    </VoahFrameWrapper>
  );
}
