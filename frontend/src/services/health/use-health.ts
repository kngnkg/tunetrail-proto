import useSWR from 'swr';

// APIの稼働状態を表す
export interface Health {
  health: string;
  database: string;
}

export interface UseHealth {
  health?: Health;
  /**
   * ロードフラグ
   */
  isLoading: boolean;
  /**
   * エラーフラグ
   */
  isError: boolean;
}

const useHealth = (): UseHealth => {
  const { data, error } = useSWR<Health>('http://localhost:8080/health');

  return {
    health: data,
    isLoading: !error && !data,
    isError: !!error,
  };
};

export default useHealth;
