import useSWR from 'swr';
jest.mock('swr');

import { renderHook } from '@testing-library/react';
import useHealth, { Health } from './use-health';

describe('useHealth', () => {
  it('fetchが成功した場合useHealthがdataを返す。エラーがfalseになる。', async () => {
    // レスポンスのモック
    const mockResponse: Health = { health: 'green', database: 'green' };
    (useSWR as jest.Mock).mockReturnValueOnce({
      data: mockResponse,
      error: null,
    });

    const { result } = renderHook(() => useHealth());

    expect(result.current.isLoading).toBe(false);
    expect(result.current.isError).toBe(false);

    const expectedHealth: Health = { health: 'green', database: 'green' };
    expect(result.current.health).toEqual(expectedHealth);
  });

  it('fetchが失敗した場合、isLoadingがfalseで、isErrorがtrueになる', () => {
    // エラーのモック
    const mockError = new Error('Request failed');
    (useSWR as jest.Mock).mockReturnValueOnce({
      data: null,
      error: mockError,
    });

    const { result } = renderHook(() => useHealth());

    expect(result.current.isLoading).toBe(false);
    expect(result.current.isError).toBe(true);
  });
});
