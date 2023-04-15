import { ApiContext } from '@/types/apiContext';
import { Health } from '@/types/health';
import { fetcher } from '@/utils';
import useSWR from 'swr';

export interface UseHealth {
  health?: Health;
  // ロードフラグ
  isLoading: boolean;
  // エラーフラグ
  isError: boolean;
}

const useHealth = (context: ApiContext): UseHealth => {
  const { data, error } = useSWR<Health>(`${context.apiRoot}/health`, fetcher);

  return {
    health: data,
    isLoading: !error && !data,
    isError: !!error,
  };
};

export default useHealth;
