/* eslint-disable @typescript-eslint/no-explicit-any */
import { pageStore } from '../../stores/pageStore';
import gql from 'graphql-tag';

export default class PersonasService {
  async conductores() {
    const store = pageStore();
    const sql = gql`
      query conductores{
        conductores{
          id
          nombres
          apellidos
          username
          foto_url
          telefono
          documento
          fecha_nac
          domicilio
          correo
          registrado
          estado
          dataname
          roles{id nombre}
        }
      }
    `;

    const res: any = await store.run_graphql(
      sql,
      { },
      { showNotificacion: false, showNotificacionError: true }
    );
    return res;
  }

  async clientes() {
    const store = pageStore();
    const sql = gql`
      query clientes{
        clientes{
          id
          nombres
          apellidos
          username
          foto_url
          telefono
          documento
          fecha_nac
          domicilio
          correo
          registrado
          estado
          dataname
          roles{id nombre}
        }
      }
    `;

    const res: any = await store.run_graphql(
      sql,
      { },
      { showNotificacion: false, showNotificacionError: true }
    );
    return res;
  }
}
