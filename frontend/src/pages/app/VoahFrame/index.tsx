import { useAtom } from 'jotai';
import { SendIcon } from 'lucide-react';
import { useEffect, useRef, useState } from 'react';

import { userAtom } from '@/atom';
import { css } from '@/stitches.config';

import { useVoahMessageFunc } from './logic';
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

  const voahMessages = useVoahMessageFunc(port1);

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
          voahMessages.frame.initDone(url);
          break;
        case 'VOAH__USER_GET_TOKEN':
          voahMessages.user.getToken(user.accessToken);
          break;
        case 'VOAH__SIDEBAR_SET_INFO':
          voahMessages.sidebar.setSidebarInfo(
            e.data.data as {
              title: string;
              desc: string;
              hideDesc: boolean;
            },
          );
          break;
        case 'VOAH__SIDEBAR_SET_MENU':
          voahMessages.sidebar.setSidebarMenu(
            e.data.data as {
              categories: Array<{
                id: string;
                name: string;
              }>;
              menus: Array<{
                icon: string;
                name: string;
                onClick: string;
                mentioned?: number;
                isFocused?: boolean;
                categoryId?: string;
                subButton?: {
                  icon: string;
                  onClick: string;
                };
                subMenu?: Array<{
                  icon: string;
                  name: string;
                  onClick: string;
                }>;
              }>;
            },
          );
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
          <SendIcon size={20} />
        </AddressBarBtn>
      </AddressBarWrapper>
      <iframe
        className={css({
          width: '100%',
          height: '100%',
          border: '2px solid $gray300',
          borderRadius: '0.75rem',
        })()}
        ref={frameRef}
        title="module"
        src={url}></iframe>
    </VoahFrameWrapper>
  );
}
