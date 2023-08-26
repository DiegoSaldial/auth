/* eslint-disable @typescript-eslint/no-explicit-any */
import { DocumentNode } from 'graphql/language/ast';
import { Opciones, InsOpciones } from '../../types/responses';
import { useMutation } from '@vue/apollo-composable';
import { watch } from 'vue';
import { evaluarError } from './onerrorgraphql';
import { mostrarNotificacion } from './oncomplete';

export const run_graphql_mutation = (
  sql: DocumentNode,
  variables = {},
  opciones: Opciones = InsOpciones
) => {
  return new Promise<any>((resolve, reject) => {
    const {
      mutate: envio,
      error,
      onDone,
    } = useMutation(sql, {
      variables: { ...variables },
    });

    watch(envio, (value) => {
      console.log('mutate:', value);
    });
    onDone(({ data }) => {
      mostrarNotificacion(opciones);
      resolve(data);
    });

    watch(error, (value) => {
      if (opciones.showNotificacionError) {
        evaluarError(error, value);
      }
      reject(value);
    });
  });
};
