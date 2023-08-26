/* eslint-disable @typescript-eslint/no-explicit-any */
import { ApolloError } from '@apollo/client/core/index.js';
import { Ref } from 'vue-demi';
import { ObjetoErr } from '../../types/responses';
import { notificar } from './notificacion';

export function evaluarError(
  error: Ref<ApolloError | null>,
  value: ApolloError | null
) {
  const men = value?.message;
  const net: any = error.value?.networkError;
  const resultado = net?.result;
  let gqlgen_msg = '';

  if (typeof resultado == 'string') {
    gqlgen_msg = resultado;
  } else if (!resultado) {
    gqlgen_msg = '';
  } else {
    const errores = resultado.errors;
    if (errores) {
      let txterr = '';
      console.log('error', errores);
      errores.forEach((e: ObjetoErr) => {
        txterr += `${e.message}, `;
        let per = '';
        if (e.path) {
          e.path.forEach((p) => (per += `${p}, `));
        }
        if (per) txterr += `[${per}]`;
      });
      gqlgen_msg = txterr;
    } else {
      gqlgen_msg = men ? men : '';
    }
  }

  notificar('' + men, gqlgen_msg, 'primary');
}
