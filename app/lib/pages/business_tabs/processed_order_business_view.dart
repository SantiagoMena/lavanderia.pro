import 'package:flutter/material.dart';
import 'package:flutter_gen/gen_l10n/app_localizations.dart';
import 'package:lavanderiapro/auth/register_business.dart';
import 'package:lavanderiapro/services/get_profile_service.dart';

class ProcessedOrderBusiness extends StatefulWidget {
   const ProcessedOrderBusiness({super.key});

  @override
  State<ProcessedOrderBusiness> createState() => _ProcessedOrderBusinessState();
}

class _ProcessedOrderBusinessState extends State<ProcessedOrderBusiness> {

  @override
  Widget build(BuildContext context) {

    return Scaffold(
        appBar: AppBar(
        title: Text("Order from Client Name"),
      ),
      body: LayoutBuilder(
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
                    child: Text("🧾", style: TextStyle(fontSize: 70),)
                  ),
                  Padding(padding: EdgeInsets.only(bottom: MediaQuery.of(context).viewInsets.bottom))
                ],
               ),
              ),
        );
      }
      ),
    );
  }
}
