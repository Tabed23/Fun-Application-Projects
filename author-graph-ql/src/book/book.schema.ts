import { Field, ObjectType, Int, Float } from "@nestjs/graphql";

@ObjectType()
export class Book {
    @Field((type) => Int)
    id: number;

    @Field()
    title: string;

    @Field((type)=> Int)
    price: number
}
