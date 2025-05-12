import 'package:flutter/material.dart';
import 'package:pet_rock_flutter/config/supabase_config.dart';
import 'app.dart';

void main() async {
  WidgetsFlutterBinding.ensureInitialized();
  await SupabaseConfig.init();
  runApp(const PetRockApp());
}
