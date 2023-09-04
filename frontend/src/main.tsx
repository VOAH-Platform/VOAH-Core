import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import { Provider as AlertProvider, positions, transitions } from 'react-alert';
import { createRoot } from 'react-dom/client';

import { AlertTemplate } from '@/components/AlertTemplate/index.tsx';

import App from './App.tsx';

import '@/assets/fonts/SUIT.css';

const queryClient = new QueryClient();

const alertOptions = {
  position: positions.TOP_CENTER,
  timeout: 5000,
  offset: '6px 10px',
  transition: transitions.FADE,
};

createRoot(document.getElementById('root')!).render(
  <QueryClientProvider client={queryClient}>
    <AlertProvider template={AlertTemplate} {...alertOptions}>
      <App />
    </AlertProvider>
  </QueryClientProvider>,
);
