'use server';
import { AuthError } from 'next-auth';
import type { AuthProvider } from '@toolpad/core';
import { signIn } from '../../../auth';

export default async function serverSignIn(provider: AuthProvider, formData: FormData, callbackUrl?: string) {
  try {
    return await signIn(provider.id, {      
    ...(formData && { email: formData.get('email'), password: formData.get('password') }),      
      redirectTo: callbackUrl ?? '/',
    });
  } catch (error) {
    if (error instanceof Error && error.message === 'NEXT_REDIRECT') {
                
      throw error;
    }
    if (error instanceof AuthError) {
      return {
        error: error.type === 'CredentialsSignin' ? 'Invalid credentials.' : 'An error with Auth.js occurred.',
        type: error.type,
      };
    }
    return {
      error: 'Something went wrong.',
      type: 'UnknownError',
    };
  }
}