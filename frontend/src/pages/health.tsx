import { NextPage } from 'next';
import HealthCheckContainer from '../../HealthCheckContainer';

const HealthCheckPage: NextPage = () => {
  return (
    <>
      <p>Health Check Page</p>
      <HealthCheckContainer />
    </>
  );
};

export default HealthCheckPage;
