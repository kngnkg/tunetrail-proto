import { ApiContext, Health } from "@/types"
import useSWR from "swr"

import { fetcher } from "@/lib/fetcher"

export interface UseHealth {
  health?: Health
  // ロードフラグ
  isLoading: boolean
  // エラーフラグ
  isError: boolean
}

const useHealth = (context: ApiContext): UseHealth => {
  const { data, error } = useSWR<Health>(`${context.apiRoot}/health`, fetcher)

  return {
    health: data,
    isLoading: !error && !data,
    isError: !!error,
  }
}

export default useHealth
