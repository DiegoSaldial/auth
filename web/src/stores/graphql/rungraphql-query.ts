/* eslint-disable @typescript-eslint/no-explicit-any */
import { DocumentNode } from 'graphql/language/ast';
import { Opciones, InsOpciones } from '../../types/responses';
import { useQuery } from '@vue/apollo-composable';
import { watch } from 'vue';
import { evaluarError } from './onerrorgraphql';

export const run_graphql_query = (
  sql: DocumentNode,
  variables = {},
  opciones: Opciones = InsOpciones
) => {
  return new Promise<any>((resolve, reject) => {
    const { error, result } = useQuery(sql, variables, {
      fetchPolicy: 'no-cache',
    });

    watch(result, (e) => {
      resolve(e);
    });

    watch(error, (value) => {
      if (opciones.showNotificacionError) {
        evaluarError(error, value);
      }
      reject(value);
    });
  });
};
