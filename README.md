# limit

Components: 1.Server (new_server)
            2.Client (client)
            3.Protocol buffer (server)
            4.Database(Bolt DB) (my.bolt)

Functionality: There a Client Which has a CLI to Interact. It's Perform Four types of Action.

              Command Fromat: In new_server Directory go run main.go --action="action_name" --input="input"
                
                 Action               action_name            input
                 Store                store                  Item to be store
                 Get Item             getid                  id of an item
                 List                 list                   number of item
                 Remove               rm                     id of item

There is need to run Server and Client Separately.
            
