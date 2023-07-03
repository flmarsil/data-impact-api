NAME		=	matcha

all			: 	$(NAME)

$(NAME)		:
				docker-compose -f srcs/docker-compose.yml up --force-recreate --build

install		:	
				docker-compose -f srcs/docker-compose.yml up --force-recreate --build

clean		:	
				docker-compose -f srcs/docker-compose.yml down --volumes --remove-orphans

fclean		:	clean
				docker system prune -a

.PHONY		:	all clean fclean
