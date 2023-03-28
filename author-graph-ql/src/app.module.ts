import { AppResolver } from './app.resolver';
import { Module } from '@nestjs/common';
import { GraphQLModule } from '@nestjs/graphql';
import { ApolloDriver, ApolloDriverConfig } from '@nestjs/apollo';
import { BookModule } from './book/book.module';
import { join } from 'path';
import { TypeOrmModule } from '@nestjs/typeorm';

@Module({
  imports: [
    GraphQLModule.forRoot<ApolloDriverConfig>({
      driver: ApolloDriver,
      playground: true,
      //typePaths: ['./**/*.graphql'],
      autoSchemaFile: join(process.cwd(), 'src/schema.graphql'),
      definitions: {
        path: join(process.cwd(), 'src/graphql.ts')
      },
    }),

    TypeOrmModule.forRoot(
      {
        type:'postgres',
        host: 'localhost',
        port: 5432,
        username: 'postgres',
        password: 'postgres',
        database: 'books_db',
        entities: [__dirname + '/**/*.entity{.ts,.js}'],
        synchronize: true,
      }
    ),

    BookModule
  ],
  controllers: [],
  providers: [AppResolver],
})
export class AppModule {}
