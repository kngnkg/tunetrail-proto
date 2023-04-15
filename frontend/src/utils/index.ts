export const fetcher = async (
  resource: RequestInfo,
  init?: RequestInit,
): Promise<any> => {
  const response = await fetch(resource, init);
  if (!response.ok) {
    const errorResponse = await response.json();
    const error = new Error(
      errorResponse.message ?? 'An error occurred while fetching the data.',
    );
    throw error;
  }

  return response.json();
};
