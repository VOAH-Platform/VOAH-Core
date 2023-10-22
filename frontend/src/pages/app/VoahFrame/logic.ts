export function useVoahMessageFunc(port1: MessagePort) {
  return {
    frame: {
      initDone: (url: string) => {
        console.log(`${url} is loaded!`);
      },
    },
    user: {
      getToken: (accessToken: string) => {
        port1.postMessage({
          type: 'VOAH__USER_GET_TOKEN_DONE',
          data: accessToken,
        });
      },
    },
  };
}
