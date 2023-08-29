create table `usuarios`(
    `id` integer unsigned not null auto_increment primary key, 
    `nombres` varchar(30) not null,
    `apellidos` varchar(30) not null,
    `username` varchar(25) not null unique,
    `password` varchar(100) not null,
    `telefono` varchar(20) not null,
    `documento` varchar(30) not null,
    `correo` varchar(100),
    `foto_url` varchar(200),
    `fecha_nac` date,
    `domicilio` varchar(100),
    `registrado` datetime not null default CONVERT_TZ(NOW(), @@session.time_zone, '-4:00'),
    `estado` tinyint(1) not null default 1
);

create table `roles`(
    `id` integer unsigned not null auto_increment primary key, 
    `nombre` varchar(30) not null,
    `bit` integer unsigned not null, -- serial incremental
    CONSTRAINT `unique_rol_app` UNIQUE (`nombre`, `bit`)
);

create table `rol_usuario`(
    `usuario_id` integer unsigned not null,
    `rol_id` integer unsigned not null,
    `registrado` datetime not null default CONVERT_TZ(NOW(), @@session.time_zone, '-4:00'),
    foreign key(`usuario_id`) references `usuarios`(`id`),
    foreign key(`rol_id`) references `roles`(`id`),
    primary key(`usuario_id`, `rol_id`)
);

create table `rol_permiso`(
    `metodo` varchar(30) not null,
    `descripcion` varchar(200) not null,
    `rol_bits` integer unsigned not null, 
    primary key(`metodo`,`rol_bits`)
);

create table `tokens`(
    `username` varchar(25) not null,
    `token` varchar(500) not null,
    `registrado` datetime not null default CONVERT_TZ(NOW(), @@session.time_zone, '-4:00') ON UPDATE CURRENT_TIMESTAMP, 
    primary key(`username`)
);

-- =====================================================================================================================
-- =====================================================================================================================
-- =====================================================================================================================

create table `direcciones`(
    `id` integer unsigned not null auto_increment primary key, 
    `nombre` varchar(100) not null,
    `descripcion` varchar(250) not null, 
    `ubicacion` POINT NOT NULL,
    `usuario_id` integer unsigned not null,
    `registrado` datetime not null default CONVERT_TZ(NOW(), @@session.time_zone, '-4:00'),
    foreign key(`usuario_id`) references `usuarios`(`id`)
);

create table `categoria_vehiculos`(
    `id` integer unsigned not null auto_increment primary key, 
    `nombre` varchar(30) not null UNIQUE,
    `descripcion` varchar(250) not null
);

create table `vehiculos`(
    `id` integer unsigned not null auto_increment primary key, 
    `placa` varchar(15) not null UNIQUE,
    `puertas` smallint not null,
    `capacidad` smallint not null,
    `descripcion` varchar(250) not null,
    `color` varchar(30) not null, 
    `modelo` varchar(30) not null, 
    `anio` smallint not null,
    `foto_url` varchar(200),
    `categoria_id` integer unsigned not null, 
    `estado` tinyint(1) not null default 1,
    `registrado` datetime not null default CONVERT_TZ(NOW(), @@session.time_zone, '-4:00'),
    foreign key(`categoria_id`) references `categoria_vehiculos`(`id`)
);

create table `conductor_vehiculos`(
    `usuario_id` integer unsigned not null,
    `vehiculo_id` integer unsigned not null,
    `estado` tinyint(1) not null default 1,
    `registrado` datetime not null default CONVERT_TZ(NOW(), @@session.time_zone, '-4:00'),
    foreign key(`usuario_id`) references `usuarios`(`id`),
    foreign key(`vehiculo_id`) references `vehiculos`(`id`),
    primary key(`usuario_id`, `vehiculo_id`)
);

create table `viajes`(
    `id` integer unsigned not null auto_increment primary key, 
    `pasajero_id` integer unsigned not null,
    `conductor_id` integer unsigned,
    `estado` tinyint(1) not null default 1, 
    `descripcion` varchar(250) not null, 
    `origen` POINT NOT NULL,
    `destino` POINT,
    `categoria_id` integer unsigned not null, 
    `registrado` datetime not null default CONVERT_TZ(NOW(), @@session.time_zone, '-4:00'),
    foreign key(`pasajero_id`) references `usuarios`(`id`),
    foreign key(`conductor_id`) references `conductor_vehiculos`(`usuario_id`),
    foreign key(`categoria_id`) references `categoria_vehiculos`(`id`) 
);

create table `viajes_locations`(
    `id` integer unsigned not null auto_increment primary key, 
    `origen` POINT NOT NULL,
    `viaje_id` integer unsigned not null,
    `registrado` datetime not null default CONVERT_TZ(NOW(), @@session.time_zone, '-4:00'),
    foreign key(`viaje_id`) references `viajes`(`id`)
);





-- =====================================================================================================================
-- =====================================================================================================================
-- =====================================================================================================================

insert into usuarios(nombres,apellidos,username,password,telefono,documento) values('Usuario','Administrador','admin','admin','+591 78227092','10721310');
insert into roles(nombre,bit) values('administrador',1);
insert into roles(nombre,bit) values('pasajero',2);
insert into roles(nombre,bit) values('conductor',3);
insert into rol_usuario(usuario_id,rol_id) values(1,1);

-- administrador    = 1
-- invitado         = 2
-- pasajero         = 3
-- otros            = 4
-- se asigna la suma de valores de roles permitidos
-- valorAsignar = 1 << (rol_bit-1) 
-- se verifica: rol_bits & 4 == 4 para el rol 4
-- se verifica: 5 & 2 == 2 para el rol 2
insert into rol_permiso(rol_bits,metodo,descripcion) values
    (5,"usuarioByUsername","Recupera datos de un usuario por su username"),
    (1,"permisos","Recupera la lista de todos los permisos disponibles"),  
    (1,"roles","Recupera la lista de todos los roles disponibles"),  
    (1,"usuarios","Recupera la lista de todos los usuarios registrados"),  
    (1,"permisosByRol","Recupera la lista de todos los permisos disponibles por rol"),
    (1,"createUsuario","Registra un nuevo usuario"), 
    (1,"updateUsuario","Actualiza los datos de un usuario"), 
    (1,"createRol","Registra un nuevo Rol"),
    (1,"deleteRol","Elimina un rol del sistema"),
    (1,"modificarRol","Actualiza los datos de un rol");

-- app
insert into rol_permiso(rol_bits,metodo,descripcion) values
    (1,"direccionesByUsuario","Recupera datos de un usuario por su username"),
    (1,"categoria_vehiculos","Recupera la lista de todos los permisos disponibles"),  
    (1,"categoria_vehiculo","Recupera la lista de todos los roles disponibles"),  
    (1,"vehiculos","Recupera la lista de todos los usuarios registrados"),  
    (1,"vehiculo","Recupera la lista de todos los permisos disponibles por rol"),
    (1,"vehiculosByCategoria","Registra un nuevo usuario"), 
    (1,"vehiculosByConductor","Actualiza los datos de un usuario"), 
    (1,"conductoresByVehiculo","Registra un nuevo Rol"),
    (1,"conductores","Elimina un rol del sistema"),
    (1,"clientes","Elimina un rol del sistema"),
    (1,"viajesCercanos","Elimina un rol del sistema"),
    (1,"viajesByUsuario","Elimina un rol del sistema"),
    (1,"viajesByPasajero","Elimina un rol del sistema"),
    (1,"viajesByConductor","Elimina un rol del sistema"),
    (1,"viajesByCategoria","Elimina un rol del sistema"),
    (1,"createUpdateDireccion","Elimina un rol del sistema"),
    (1,"createUpdateCategoriaVehiculo","Elimina un rol del sistema"),
    (1,"createUpdateVehiculo","Elimina un rol del sistema"),
    (1,"asignarVehiculo","Elimina un rol del sistema"),
    (1,"quitarVehiculo","Elimina un rol del sistema"),
    (1,"createViaje","Elimina un rol del sistema"),
    (1,"aceptarViaje","Elimina un rol del sistema"),
    (1,"cancelarViaje","Elimina un rol del sistema"),
    (1,"finalizarViaje","Elimina un rol del sistema"),
    (1,"setUbicacionViaje","Elimina un rol del sistema");


