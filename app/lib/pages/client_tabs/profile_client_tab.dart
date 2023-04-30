import 'package:flutter/material.dart';
import 'package:flutter_gen/gen_l10n/app_localizations.dart';
import 'package:lavanderiapro/auth/register_business.dart';
import 'package:lavanderiapro/services/get_profile_service.dart';

class ProfileClientTab extends StatefulWidget {
   const ProfileClientTab({super.key, this.token});

  final String? token;

  @override
  State<ProfileClientTab> createState() => _ProfileClientTabState();
}

class _ProfileClientTabState extends State<ProfileClientTab> {
  final _formNameKey = GlobalKey<FormState>();
  final _formPasswordKey = GlobalKey<FormState>();

  TextEditingController nameController = TextEditingController();
  TextEditingController passwordController = TextEditingController();
  TextEditingController newPasswordController = TextEditingController();

  @override
  Widget build(BuildContext context) {

    Profile? profile;

    return LayoutBuilder(
      builder: (BuildContext context, BoxConstraints viewportConstraints) {
      return Align(
        alignment: Alignment.topCenter,
        child: SingleChildScrollView(
              reverse: true,
              child: Column(
                crossAxisAlignment: CrossAxisAlignment.start,
                mainAxisAlignment: MainAxisAlignment.start,
                mainAxisSize: MainAxisSize.min,
                children: [
                  Padding(
                    padding: const EdgeInsets.symmetric(horizontal: 8, vertical: 16),
                    child: Center(
                      child: ElevatedButton(
                        style: ElevatedButton.styleFrom(
                          minimumSize: const Size.fromHeight(50),
                          backgroundColor: Colors.green,
                        ),
                        onPressed: () {

                        },
                        child: Text(AppLocalizations.of(context)!.manageAddressesButtonLabel),
                      ),
                    ),
                  ),
                  Form(
                       key: _formNameKey,
                       child: Column(
                         children:[
                           Padding(
                             padding: const EdgeInsets.symmetric(horizontal: 8, vertical: 16),
                             child: TextFormField(
                               controller: nameController,
                               decoration: const InputDecoration(border: OutlineInputBorder(), label: ClientNameLabel()),
                               validator: (value) {
                                 if(value == null || value.isEmpty){
                                   return AppLocalizations.of(context)!.emptyNameAlert;
                                 }
                                 return null;
                               },
                             ),
                           ),
                           Padding(
                             padding: const EdgeInsets.symmetric(horizontal: 8, vertical: 16),
                             child: Center(
                               child: ElevatedButton(
                                 style: ElevatedButton.styleFrom(
                                   minimumSize: const Size.fromHeight(50),
                                   backgroundColor: Colors.green,
                                 ),
                                 onPressed: () {
                                   if(_formNameKey.currentState!.validate()){
                                     // Change Name

                                   } else {
                                     ScaffoldMessenger.of(context).showSnackBar(
                                         const SnackBar(content: FillInputSnackBar())
                                     );
                                   }
                                 },
                                 child: Text(AppLocalizations.of(context)!.changeNameButtonLabel),
                               ),
                             ),
                           ),
                         ],
                       ),
                     ),
                  Form(
                        key: _formPasswordKey,
                         child: Column(
                             children:[
                           Padding(
                             padding: const EdgeInsets.symmetric(horizontal: 8, vertical: 16),
                             child: TextFormField(
                               controller: passwordController,
                               obscureText: true,
                               decoration: const InputDecoration(border: OutlineInputBorder(), label: CurrentPasswordLabel()),
                               validator: (value) {
                                 if(value == null || value.isEmpty){
                                   return AppLocalizations.of(context)!.emptyPasswordAlert;
                                 }
                                 return null;
                               },
                             ),
                           ),
                           Padding(
                             padding: const EdgeInsets.symmetric(horizontal: 8, vertical: 16),
                             child: TextFormField(
                               controller: newPasswordController,
                               obscureText: true,
                               decoration: const InputDecoration(border: OutlineInputBorder(), label: NewPasswordLabel()),
                               validator: (value) {
                                 if(value == null || value.isEmpty){
                                   return AppLocalizations.of(context)!.emptyPasswordAlert;
                                 }
                                 return null;
                               },
                             ),
                           ),
                           Padding(
                             padding: const EdgeInsets.symmetric(horizontal: 8, vertical: 16),
                             child: Center(
                               child: ElevatedButton(
                                 style: ElevatedButton.styleFrom(
                                   minimumSize: const Size.fromHeight(50),
                                   backgroundColor: Colors.green,
                                 ),
                                 onPressed: () {
                                   if(_formPasswordKey.currentState!.validate()){
                                     // Change Name

                                   } else {
                                     ScaffoldMessenger.of(context).showSnackBar(
                                         const SnackBar(content: FillInputSnackBar())
                                     );
                                   }
                                 },
                                 child: Text(AppLocalizations.of(context)!.changePasswordButtonLabel),
                               ),
                             ),
                           ),
                           ],
                         ),
                       ),
                  Padding(
                   padding: const EdgeInsets.symmetric(horizontal: 8, vertical: 16),
                   child: Center(
                     child: ElevatedButton(
                       style: ElevatedButton.styleFrom(
                         minimumSize: const Size.fromHeight(50),
                         backgroundColor: Colors.green,
                       ),
                       onPressed: () {
                         Navigator.pushNamedAndRemoveUntil(context, '/login', (_) => false);
                       },
                       child: Text("Logout"),
                     ),
                   ),
                 ),
                  Padding(padding: EdgeInsets.only(bottom: MediaQuery.of(context).viewInsets.bottom))
                ],
               ),
              ),
        );
      }
    );
  }
}

class NewPasswordLabel extends StatelessWidget {
  const NewPasswordLabel({
    super.key,
  });

  @override
  Widget build(BuildContext context) {
    return Text(AppLocalizations.of(context)!.newPasswordLabel);
  }
}

class CurrentPasswordLabel extends StatelessWidget {
  const CurrentPasswordLabel({
    super.key,
  });

  @override
  Widget build(BuildContext context) {
    return Text(AppLocalizations.of(context)!.currentPasswordLabel);
  }
}

class ClientNameLabel extends StatelessWidget {
  const ClientNameLabel({
    super.key,
  });

  @override
  Widget build(BuildContext context) {
    return Text(AppLocalizations.of(context)!.clientNameLabel);
  }
}

class FillInputSnackBar extends StatelessWidget {
  const FillInputSnackBar({
    super.key,
  });

  @override
  Widget build(BuildContext context) {
    return Text(AppLocalizations.of(context)!.snackBarFillInput);
  }
}

