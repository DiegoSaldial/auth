/* eslint-disable @typescript-eslint/no-explicit-any */
import { pageStore } from '../../stores/pageStore';
import gql from 'graphql-tag';
import { CreateConductorVehiculos, CreateVehiculos } from '../../types/taxis/vehiculos';

export default class VehiculosService {
  async vehiculos() {
    const store = pageStore();
    const sql = gql`
      query{
        vehiculos{
          id
          placa
          puertas
          capacidad
          descripcion
          color
          modelo
          anio
          categoria_id
          foto_url
          estado
          registrado
          conductor_id
          conductor
          categoria
          estado_conductor_vehiculo
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

  async vehiculosByConductor(usuario_id:string) {
    const store = pageStore();
    const sql = gql`
      query vehiculosByConductor($usuario_id:ID!){
        vehiculosByConductor(usuario_id:$usuario_id){
          id
          placa
          puertas
          capacidad
          descripcion
          color
          modelo
          anio
          categoria_id
          foto_url
          estado
          registrado
          conductor_id
          conductor
          categoria
          estado_conductor_vehiculo
        }
      }
    `;

    const res: any = await store.run_graphql(
      sql,
      { usuario_id:usuario_id },
      { showNotificacion: false, showNotificacionError: true }
    );
    return res;
  }

  async vehiculosByCategoria(categoria_id:string) {
    const store = pageStore();
    const sql = gql`
      query vehiculosByCategoria($categoria_id:ID!){
        vehiculosByCategoria(categoria_id:$categoria_id){
          id
          placa
          puertas
          capacidad
          descripcion
          color
          modelo
          anio
          categoria_id
          foto_url
          estado
          registrado
          conductor_id
          conductor
          categoria
          estado_conductor_vehiculo
        }
      }
    `;

    const res: any = await store.run_graphql(
      sql,
      { categoria_id:categoria_id },
      { showNotificacion: false, showNotificacionError: true }
    );
    return res;
  }

  async createUpdateVehiculo(input:CreateVehiculos) {
    const store = pageStore();
    const sql = gql`
      mutation createUpdateVehiculo($input:CreateVehiculos!){
        createUpdateVehiculo(input:$input){id}
      }
    `;

    const res: any = await store.run_graphql(
      sql,
      { input:input },
      { showNotificacion: false, showNotificacionError: true }
    );
    return res;
  }

  async vehiculo(id:string) {
    const store = pageStore();
    const sql = gql`
      query vehiculo($id:ID!){
        vehiculo(id:$id){
          id
          placa
          puertas
          capacidad
          descripcion
          color
          modelo
          anio
          categoria_id
          foto_url
        }
      }
    `;

    const res: any = await store.run_graphql(
      sql,
      { id:id },
      { showNotificacion: false, showNotificacionError: true }
    );
    return res;
  }

  async asignarVehiculo(input:CreateConductorVehiculos) {
    const store = pageStore();
    const sql = gql`
      mutation asignarVehiculo($input:CreateConductorVehiculos!){
        asignarVehiculo(input:$input){usuario_id vehiculo_id}
      }
    `;

    const res: any = await store.run_graphql(
      sql,
      { input:input },
      { showNotificacion: false, showNotificacionError: true }
    );
    return res;
  }

  async quitarVehiculo(input:CreateConductorVehiculos) {
    const store = pageStore();
    const sql = gql`
      mutation quitarVehiculo($input:CreateConductorVehiculos!){
        quitarVehiculo(input:$input)
      }
    `;

    const res: any = await store.run_graphql(
      sql,
      { input:input },
      { showNotificacion: false, showNotificacionError: true }
    );
    return res;
  }
}
