import { Query } from '@nestjs/graphql';
import { Resolver } from '@nestjs/graphql';

@Resolver(of => String)
export class AppResolver {

    @Query(returns =>  String)
    index(): String {
        return 'NestJs is up and running';
    }
}
