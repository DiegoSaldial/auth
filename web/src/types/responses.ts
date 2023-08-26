/* eslint-disable @typescript-eslint/no-explicit-any */
import { MutateResult } from '@vue/apollo-composable';

export type CallbackReturn = (value: [] | MutateResult<any>) => void;
export type CallbackReturnRow = (value: MutateResult<any> | any) => void;

export interface ObjetoErr {
  message: string;
  path: [];
}

export interface Opciones {
  showNotificacion?: boolean;
  showNotificacionError?: boolean;
  // responseNullOnError?: boolean;
}

export const InsOpciones: Opciones = {
  showNotificacion: true,
  showNotificacionError: true,
  // responseNullOnError: true,
};

export interface Site {
  logo: string;
  ratio: string;
  name: string;
}

export const InstanciaSite: Site = {
  logo: '',
  ratio: '',
  name: '',
};
