const nextJest = require('next/jest');
const createJestConfig = nextJest({ dir: './' });
const customJestConfig = {
  testPathIgnorePatterns: ['<rootDir>/.next/', '<rootDir>/node_modules/'],
  setupFilesAfterEnv: ['<rootDir>/jest.setup.js'],
  moduleDirectories: ['node_modules', '<rootDir>/src'],
  testEnvironment: 'jsdom',
  transform: {
    '^.+\\.(js|jsx|ts|tsx)$': 'ts-jest',
  },
  moduleNameMapper: {
    '\\.(css|less|scss|sass)$': 'identity-obj-proxy',
  },
};
module.exports = createJestConfig(customJestConfig);
