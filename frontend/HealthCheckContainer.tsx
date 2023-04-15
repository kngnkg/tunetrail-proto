import useHealth from '@/services/health/use-health';
import { ApiContext } from '@/types/apiContext';

const context: ApiContext = {
  apiRoot: process.env.NEXT_PUBLIC_API_BASE_PATH || '/api/proxy',
};

const HealthCheckContainer = () => {
  const { health, isLoading, isError } = useHealth(context);
  return (
    <>
      <p>health: {health?.health}</p>
      <p>database: {health?.database}</p>
      {isLoading && <p>Loading...</p>}
      {isError && <p>Error</p>}
    </>
  );
};

export default HealthCheckContainer;
