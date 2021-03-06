import { render, fireEvent } from '@testing-library/svelte';
import { get } from 'svelte/store';

import Register from './Register.svelte';
import { jwt, notifications } from '../stores';

describe('Register page', () => {
  it('has a okno heading', () => {
    const { getByText } = render(Register);
    const heading = getByText('okno');
    expect(heading).toBeInstanceOf(HTMLHeadingElement);
  });
  it('sends the log in form', async () => {
    // Mock the fetch response
    global.fetch = jest.fn().mockImplementation(async () => ({
      ok: true,
      json: async () => ({ token: 'test-token' }),
    }));
    // Render and interact with DOM
    const { getByLabelText, getByText } = render(Register);
    const userInput = getByLabelText('User');
    const passwordInput = getByLabelText('Password');
    await Promise.all([
      fireEvent.input(userInput, { target: { value: 'my-user' } }),
      fireEvent.input(passwordInput, { target: { value: 'my-password' } }),
    ]);
    const button = getByText('Register');
    await fireEvent.click(button);
    // Assert fetch call
    expect(global.fetch).toHaveBeenCalledTimes(1);
    expect(global.fetch).toHaveBeenCalledWith(':4433/register', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({
        user: 'my-user',
        password: 'my-password',
      }),
    });
    // Check if jwt in the state was updated
    expect(get(jwt)).toBe('test-token');
    // Check if the user was notified
    expect(
      get(notifications).find((n) => n.message === 'Could not register'),
    ).not.toBeNull();
  });
  it('notifies on register fail', async () => {
    global.fetch = jest.fn().mockImplementation(async () => ({ ok: false }));
    const { getByText } = render(Register);
    const button = getByText('Register');
    await fireEvent.click(button);
    expect(
      get(notifications).find((n) => n.message === 'Could not register'),
    ).not.toBeNull();
  });
});
