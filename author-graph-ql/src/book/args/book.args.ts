import { Field, InputType, Int } from "@nestjs/graphql";

@InputType()
export class CreateBookArgs {

    @Field()
    title: string;

    @Field((type) => Int)
    price: number;
}

@InputType()
export class UpdateBookArgs {
    @Field((type)=>Int)
    id : number;

    @Field()
    title: string;

    @Field((type) => Int)
    price: number;
}
